import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Permission } from './permission.types';

@Injectable({
  providedIn: 'root'
})
export class PermissionService {
  private apiUrl = 'http://localhost:8081/api/permissions';

  constructor(private _httpClient: HttpClient) { }

  public newPermission(permission: Permission): Observable<any> {
    return this._httpClient.post(this.apiUrl + '/new', permission);
  }

  public getPermissions(): Observable<any> {
    return this._httpClient.get(this.apiUrl + '/all');
  }

  public getPermissionByID(id: any): Observable<any> {
    return this._httpClient.get(this.apiUrl + '/' + id);
  }

  public updatePermission(id: any, permission: Permission): Observable<any> {
    return this._httpClient.put(this.apiUrl + '/' + id, permission);
  }

  public removePermission(id: any): Observable<any> {
    return this._httpClient.delete(this.apiUrl + '/' + id);
  }

}
