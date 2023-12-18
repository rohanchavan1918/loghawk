import { Injectable } from "@angular/core";

@Injectable({
  providedIn: "root",
})
export class AuthService {
  
  isLoggedIn = false;
  
  constructor() {}
  
  isAuthenticated() {
    const isLoggedIn =  localStorage.getItem('isLoggedIn')
    return isLoggedIn == '1' ? true  : false;
  }

}
