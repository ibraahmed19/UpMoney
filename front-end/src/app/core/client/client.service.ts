import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Client, SearchReq } from './client.types';

@Injectable({
  providedIn: 'root'
})
export class ClientService {
  private apiUrl = 'http://localhost:8081/api/clients';

  constructor(private _httpClient: HttpClient) { }

  public newClient(client: Client): Observable<any> {
    return this._httpClient.post(this.apiUrl + '/new', client);
  }

  public getClients(): Observable<any> {
    return this._httpClient.get(this.apiUrl + '/all');
  }

  public getClientByID(id: any): Observable<any> {
    return this._httpClient.get(this.apiUrl + '/' + id);
  }

  public searchClients(searchReq: SearchReq): Observable<any> {
    return this._httpClient.post(this.apiUrl + '/search', searchReq);
  }

  public updateClient(id: any, client: Client): Observable<any> {
    return this._httpClient.put(this.apiUrl + '/' + id, client);
  }

  public deleteClient(id: any): Observable<any> {
    return this._httpClient.delete(this.apiUrl + '/' + id);
  }

}
