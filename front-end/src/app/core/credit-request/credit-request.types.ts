export interface CreditRequest {

        id?; number,
        client_id: number,
        type: string,
        amount_requested: number,
        justification: string,
        request_status: string,
        submission_date: string,
        approval_date: string,
        rejection_date: string,
        approved_amount: number,
        interest_rate: number,
        duration: number,
        monthly_payments: number,
        notes: string
}

export interface SearchReq{

        id?; number,
        client_id?: number,
        type?: string,
        amount_requested?: number,
        justification?: string,
        request_status?: string,
        submission_date?: string,
        approval_date?: string,
        rejection_date?: string,
        approved_amount?: number,
        interest_rate?: number,
        duration?: number,
        monthly_payments?: number,
        notes?: string
}