export interface ResponseRoot<T = unknown> {
    data: T
    message: string
    success: boolean
    timestamp?: string
    errors?: string[]
}

// Interfaces específicas para casos de uso comunes
export interface SuccessResponse<T> extends Omit<ResponseRoot<T>, 'success'> {
    success: true
    data: T
}

export interface ErrorResponse extends Omit<ResponseRoot<never>, 'success' | 'data'> {
    success: false
    data?: null
    errors: string[]
}

// Type union para responses
export type ApiResponse<T> = SuccessResponse<T> | ErrorResponse