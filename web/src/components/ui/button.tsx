interface ButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
    color?: 'primary' | 'secondary' | 'accent' | 'light' | 'dark'
    size?: 'sm' | 'md' | 'lg' | 'xl'
    pill?: boolean
}

export default function Button({ color, size, pill, className, children, ...rest }: ButtonProps) {
    color = color || 'light'
    const textColor = color === 'light' ? 'dark' : 'light'

    size = size || 'md'
    const padding = {
        sm: 'px-2 py-1',
        md: 'px-3 py-1',
        lg: 'px-4 py-2',
        xl: 'px-5 py-2',
    }[size]

    pill = pill || false
    const rounded = pill ? 'rounded-full' : 'rounded'

    return (
        <button {...rest} className={`bg-${color} text-${textColor} text-${size} hover:bg-${color}/50 font-bold ${padding} ${rounded} ${className}`}>
            {children}
        </button>
    )
}