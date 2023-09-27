export interface Paymee{
        id?: number,
        amount: number,
        note: string,
        first_name: string,
        last_name: string,
        email: string,
        phone: string,
        return_url: string,
        cancel_url: string,
        webhook_url: string,
        order_id: string
}

export interface searchReq{
    id?: number,
    amount?: number,
    note?: string,
    first_name?: string,
    last_name?: string,
    email?: string,
    phone?: string,
    return_url?: string,
    cancel_url?: string,
    webhook_url?: string,
    order_id?: string
}