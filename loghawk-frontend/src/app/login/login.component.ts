import { Component, OnInit } from '@angular/core';
import { NgForm } from '@angular/forms';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {

  constructor(
    private router: Router
  ) { }

  ngOnInit() {
  }

  onLogin(form: NgForm) {
    console.log(form.value);
    // const token = this.userAuth.authUser(form.value);

    if (['Admin','admin'].includes(form.value.userName) && ['Admin','admin'].includes(form.value.password)) {
      // alert('Login Success')
      
      localStorage.setItem('isLoggedIn', '1')
      this.router.navigate(['/dashboard']);
    } else {
      alert('Login Fail')

    }
    // if (token) {
    //   localStorage.setItem('token', token.userName)
    //   // this.alertyfy.success('Login success');
    //   this.router.navigate(['/'])
    // } else {
    //   // this.alertyfy.error('Login Fail');
    //   alert('Login Fail')
    // }
  }
}
