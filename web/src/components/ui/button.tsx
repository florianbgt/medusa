interface ButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
    color: 'primary' | 'secondary' | 'accent' | 'light' | 'dark'
    size: 'sm' | 'md' | 'lg' | 'xl'
    pill?: boolean
}

export default function Button({ color, size, pill, className, children, ...rest }: ButtonProps) {
    const colorClasses = {
        primary: 'bg-primary hover:bg-primary/50 disabled:bg-primary/25 text-light',
        secondary: 'bg-secondary hover:bg-secondary/50 text-light',
        accent: 'bg-accent hover:bg-accent/50 text-light',
        light: 'bg-light hover:bg-light/50 text-dark',
        dark: 'bg-dark hover:bg-dark/50 text-light',
    }[color]
    

    const sizeClasses = {
        sm: 'text-sm px-2 py-1',
        md: 'text-md px-3 py-1',
        lg: 'text-lg px-4 py-2',
        xl: 'text-xl px-5 py-2',
    }[size]

    pill = pill || false
    const rounded = pill ? 'rounded-full' : 'rounded'

    return (
        <button {...rest} className={`${colorClasses} ${sizeClasses} ${rounded} font-bold ${className}`}>
            {children}
        </button>
    )
}