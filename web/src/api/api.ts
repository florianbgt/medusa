type Payload = {
  url: string;
  method?: string;
  data?: object;
  _retried?: true;
};

const baseUrl = "http://localhost:8080/api/";

function logout() {
  localStorage.removeItem("refresh");
  localStorage.removeItem("access");
  location.href = "/login";
}

async function refreshToken() {
  try {
    const response = await fetch(baseUrl + "token/refresh/", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ refresh: localStorage.getItem("refresh") }),
    });
    if (response.ok) {
      const { access, refresh } = await response.json();
      localStorage.setItem("refresh", refresh);
      localStorage.setItem("access", access);
    } else throw Error;
  } catch (error) {
    logout();
    throw error;
  }
};

async function onResponseError(
  originalRequest: Payload,
  response: any
): Promise<any> {
  const data = await response.json();
  if (response.status === 401) {
    await refreshToken();
    return api({ ...originalRequest, _retried: true });
  }
  throw data;
};

async function api(payload: Payload) {
  const { url, method, data, _retried } = payload;
  const accessToken = localStorage.getItem("access");

  type Headers = {
    "Content-Type": "application/json";
    Authorization?: string;
  };

  const headers: Headers = {
    "Content-Type": "application/json",
  };

  const body = data ? JSON.stringify(data) : undefined;

  if (accessToken) headers["Authorization"] = `Bearer ${accessToken}`;

  const response = await fetch(url.startsWith("http") ? url : baseUrl + url, {
    method: method || "GET",
    headers,
    body,
  });

  if (response.ok || _retried) {
    if (response.status === 204) return null;
    return await response.json();
  }

  return await onResponseError(payload, response);
};

export { api, logout };
