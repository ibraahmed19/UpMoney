import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Role, SearchReq } from './role.types';

@Injectable({
  providedIn: 'root'
})
export class RoleService {
  private apiUrl = 'http://localhost:8081/api/roles';

  constructor(private _httpClient: HttpClient) { }

  public NewRole(role: Role): Observable<any> {
    return this._httpClient.post(this.apiUrl + '/new', role);
  }

  public GetRoles(): Observable<any> {
    return this._httpClient.get(this.apiUrl + '/all');
  }

  public SearchRoles(searchReq: SearchReq): Observable<any> {
    return this._httpClient.post(this.apiUrl + '/search', searchReq);
  }

  public UpdateRole(id: any, role: Role): Observable<any> {
    return this._httpClient.put(this.apiUrl + '/' + id, role);
  }

  public DeleteRole(id: any): Observable<any> {
    return this._httpClient.delete(this.apiUrl + '/' + id);
  }
}
