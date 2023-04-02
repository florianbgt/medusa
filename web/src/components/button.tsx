interface ButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
    color?: 'primary' | 'secondary' | 'accent' | 'light' | 'dark'
    size?: 'sm' | 'md' | 'lg' | 'xl'
}

export default function Input({ color="light", size="md", className, children, ...rest }: ButtonProps) {
    const textColor = color === 'light' ? 'dark' : 'light'

    return (
        <button className={`bg-${color} hover:bg-${color}/50 text-${textColor} text-${size} hover:bg-${color}/50 font-bold px-4 py-2 rounded ${className}`}>
            {children}
        </button>
    )
}