export interface Insurance {
        id?: number,
        name: string,
        type: string,
        address: string,
        phone: string,
        email: string,
        description: string,
        managed_by: number
}

export interface SearchReq{
    id?: number,
    name?: string,
    type?: string,
    address?: string,
    phone?: string,
    email?: string,
    description?: string,
    managed_by?: number
}