interface InputProps extends React.InputHTMLAttributes<HTMLInputElement> {
    id: string
    label?: string
    elRef?: React.Ref<HTMLInputElement>
}

export default function Input({ id, label, className, elRef, ...rest }: InputProps) {
    return (
        <div className={className}>
            {label && (
                <label
                    htmlFor={id}
                    className="block mb-2 text-sm font-bold text-light"
                >
                    {label}
                </label>
            )}
            <input
                ref={elRef}
                id={id}
                {...rest}
                className="bg-light border border-primary text-dark text-sm rounded-lg focus:outline-1 focus:outline-primary block w-full p-2"
            />
        </div>
    )
}