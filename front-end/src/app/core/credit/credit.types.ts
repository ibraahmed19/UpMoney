export interface Credit{
    
        id?: number,
        client_id: number,
        type: string,
        amount: number,
        interest_rate: number,
        duration: number,
        start_date: string,
        end_date: string,
        monthly_payments: number,
        status: string,
        comments: string
}

export interface SearchReq {
        id?: number,
        client_id?: number,
        type?: string,
        amount?: number,
        interest_rate?: number,
        duration?: number,
        start_date?: string,
        end_date?: string,
        monthly_payments?: number,
        status?: string,
        comments?: string
}