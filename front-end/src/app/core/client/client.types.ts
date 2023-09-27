export interface Client{
    
        id?: number,
        id_role: number,
        role_name: string,
        first_name: string,
        last_name: string,
        adress: string,
        country: string,
        city: string,
        zip_code: number,
        phone: string,
        email: string,
        inscription_date: string,
        birth_date: string,
        gender: string,
        profession: string,
        username: string,
        password: string,
        last_login: string
}

export interface SearchReq{
    
    id?: number,
    id_role?: number,
    role_name?: string,
    first_name?: string,
    last_name?: string,
    adress?: string,
    country?: string,
    city?: string,
    zip_code?: number,
    phone?: string,
    email?: string,
    inscription_date?: string,
    birth_date?: string,
    gender?: string,
    profession?: string,
    username?: string,
    password?: string,
    last_login?: string
}