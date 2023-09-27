export interface Deposit{
    
        id?: number,
        id_transaction: number,
        iban_account: number,
        amount: number,
        date: string,
        description: string
}

export interface SearchReq{
    id?: number,
    id_transaction?: number,
    iban_account?: number,
    amount?: number,
    date?: string,
    description?: string
}