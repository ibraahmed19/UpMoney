export interface BankingAccount{
    
        id?: number,
        bank_id: number,
        client_id: number,
        iban: string,
        balance: number,
        type: string,
        openning_date: string,
        transactions: number
}

export interface SearchReq{
        id?: number,
        bank_id?: number,
        client_id?: number,
        iban?: string,
        balance?: number,
        type?: string,
        openning_date?: string,
        transactions?: number
}