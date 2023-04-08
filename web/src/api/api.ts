type Payload = {
  url: string;
  method?: "GET" | "POST" | "PUT" | "DELETE" | "PATCH";
  data?: object;
  body?: BodyInit;
  headers?: HeadersInit;
  _retried?: true;
};

export const baseUrl = process.env.NEXT_PUBLIC_BASE_API_URL || "/api";

function logout() {
  localStorage.removeItem("refresh");
  localStorage.removeItem("access");
  location.href = "/login";
}

async function refreshToken() {
  try {
    const response = await fetch(baseUrl + "/token/refresh", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ refresh: localStorage.getItem("refresh") }),
    });
    if (response.ok) {
      const { access_token, refresh_token } = await response.json();
      localStorage.setItem("access", access_token);
      localStorage.setItem("refresh", refresh_token);
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
  let { url, method, data, body, headers, _retried } = payload;
  const accessToken = localStorage.getItem("access");

  if (!headers) headers = {};

  if (accessToken) {
    headers = {
      Authorization: `Bearer ${accessToken}`,
    };
  }

  if ("Content-Type" in headers) {
    headers["Content-Type"] = "application/json"
  }

  const requestBody = data ? JSON.stringify(data) : undefined;

  url = url.startsWith("http") ? url : baseUrl + url

  const response = await fetch(url, {
    method: method || "GET",
    headers: headers,
    body: body || requestBody,
  });

  if (response.ok || _retried) {
    if (response.status === 204) return null;
    return await response.json();
  }

  return await onResponseError(payload, response);
};

export { api, logout };
