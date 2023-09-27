export interface Bank{
    
        id?: number,
        name: string,
        address: string,
        city: string,
        country: string,
        zip_code: number,
        phone: string,
        email: string,
        contact: string,
        comments: string,
        managed_by: number
}

export interface SearchReq{
    
    id?: number,
    name?: string,
    address?: string,
    city?: string,
    country?: string,
    zip_code?: number,
    phone?: string,
    email?: string,
    contact?: string,
    comments?: string,
    managed_by?: number
}