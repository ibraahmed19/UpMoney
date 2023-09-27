import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { AuthService } from 'app/core/auth/auth.service';

@Component({
  selector: 'auth-sign-up',
  templateUrl: './sign-up.component.html',
})
export class AuthSignUpComponent implements OnInit {
  signUpForm: FormGroup;
  showAlert: boolean = false;

  constructor(
    private authService: AuthService,
    private formBuilder: FormBuilder,
    private router: Router
  ) {}

  ngOnInit(): void {
    this.signUpForm = this.formBuilder.group({
      name: ['', Validators.required],
      email: ['', [Validators.required, Validators.email]],
      password: ['', Validators.required],
      role: ['user']
    });
  }

  signUp(): void {
    if (this.signUpForm.invalid) {
      return;
    }

    this.authService.signUp(this.signUpForm.value).subscribe(
      () => {
        // Navigate to the confirmation required page
        this.router.navigateByUrl('/confirmation-required');
      },
      () => {
        this.signUpForm.enable();
        this.signUpForm.reset();
        this.showAlert = true;
      }
    );
  }
}