import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { catchError, Observable, of, switchMap, throwError } from 'rxjs';
import { AuthUtils } from 'app/core/auth/auth.utils';
import { UserService } from 'app/core/user/user.service';
import { JwtHelperService } from '@auth0/angular-jwt';

@Injectable({
    providedIn: 'root'
})
export class AuthService {
    private _authenticated: boolean = false;
    private apiUrl = 'http://localhost:8081/api/auth/signup';
    private apiUrl2 = 'http://localhost:8081/api/auth/signin';

    /**
     * Constructor
     */
    constructor(
        private _httpClient: HttpClient,
        private _userService: UserService
    )
    {
    }

    // -----------------------------------------------------------------------------------------------------
    // @ Accessors
    // -----------------------------------------------------------------------------------------------------

    /**
     * Setter & getter for access token
     */
    set accessToken(token: string)
    {
        sessionStorage.setItem('accessToken', token);
    }
    set connectedUserId(userId: number){
        sessionStorage.setItem('user_id', userId.toString());
    }

    get accessToken(): string
    {
        return sessionStorage.getItem('accessToken') ?? '';
    }

    // -----------------------------------------------------------------------------------------------------
    // @ Public methods
    // -----------------------------------------------------------------------------------------------------

    /**
     * Forgot password
     *
     * @param email
     */
    forgotPassword(email: string): Observable<any>
    {
        return this._httpClient.post('api/auth/forgot-password', email);
    }

    /**
     * Reset password
     *
     * @param password
     */
    resetPassword(password: string): Observable<any>
    {
        return this._httpClient.post('api/auth/reset-password', password);
    }



    /**
     * Sign in using the access token
     */


    /**
     * Sign out
     */
    signOut(): Observable<any>
    {
        // Remove the access token from the local storage
        sessionStorage.removeItem('accessToken');
        sessionStorage.removeItem('user_id');

        // Set the authenticated flag to false
        this._authenticated = false;

        // Return the observable
        return of(true);
    }

    /**
     * Sign up
     *
     * @param user
     */
    public signUp(data: any): Observable<any> {
        return this._httpClient.post(this.apiUrl, data);
      }

    /**
     * Unlock session
     *
     * @param credentials
     */
    unlockSession(credentials: { email: string; password: string }): Observable<any>
    {
        return this._httpClient.post('api/auth/unlock-session', credentials);
    }

    /**
     * Check the authentication status
     */
// ...

    /**
     * Sign in
     *
     * @param credentials
     */
    signIn(credentials: { email: string, password: string }) {
        return this._httpClient.post<any>(`${this.apiUrl2}`, credentials).pipe(
            switchMap((response: any) => {
                // Store the access token in the local storage
                this.accessToken = response.token; // Use 'token' key to access the token in the response

                // Set the authenticated flag to true
                this._authenticated = true;

                const helper = new JwtHelperService();
                // Store the user on the user service
                const decodedToken = helper.decodeToken(response.token);
                this.connectedUserId=decodedToken.user_id;
                this._userService.user = response.user;

                // Return the response
                return of(response);
            })
        );
    }

    // ...

    /**
     * Check the authentication status
     */
    check(): Observable<boolean>
    {
        // Check if the user is logged in
        if ( this._authenticated )
        {
            return of(true);
        }

        // Check the access token availability
        if ( !this.accessToken )
        {
            return of(false);
        }

        // Check the access token expire date
        if ( AuthUtils.isTokenExpired(this.accessToken) )
        {
            return of(false);
        }

        // If the access token exists and it didn't expire, sign in using it
        return this.signInUsingToken();
    }

    /**
     * Sign in using the access token
     */
    signInUsingToken(): Observable<boolean>
    {
        // Sign in using the token
        return this._httpClient.post('api/auth/sign-in-with-token', {
            accessToken: this.accessToken
        }).pipe(
            catchError(() => {
                // Remove the invalid access token from the local storage
                sessionStorage.removeItem('accessToken');

                // Set the authenticated flag to false
                this._authenticated = false;

                // Return false
                return of(false);
            }),
            switchMap((response: any) => {
                // Replace the access token with the new one if it's available on the response object.
                if ( response.accessToken )
                {
                    this.accessToken = response.accessToken;
                }

                // Set the authenticated flag to true
                this._authenticated = true;

                // Store the user on the user service
                this._userService.user = response.user;

                // Return true
                return of(true);
            })
        );
    }
}
