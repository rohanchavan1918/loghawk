import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, NgForm, Validators } from '@angular/forms';
import { ApiServicesService } from '../services/api-services.service';
import { ActivatedRoute } from '@angular/router';
import { ToastrService } from 'ngx-toastr';


@Component({
  selector: 'app-tag-form',
  templateUrl: './tag-form.component.html',
  styleUrls: ['./tag-form.component.css']
})
export class TagFormComponent implements OnInit {

  tagID: number = null;
  addTagForm: FormGroup;

  constructor(
    private apiService: ApiServicesService,
    private route: ActivatedRoute,
    private fb: FormBuilder,
    private toastr: ToastrService
    ) { }

  ngOnInit() {
    this.initializeForm()
    this.route.params.subscribe(params => {
      console.log(params['id'],'this.tagID');
      this.tagID = params['id'] ? +params['id'] : null;
      if(![undefined, null, NaN].includes(params['id'])) {
        this.getTagById(+params['id'])    
      }
    })
  }

  initializeForm() {
    this.addTagForm = this.fb.group({
      name: [null, Validators.required],
      description: [null, Validators.required],
      SlackUrl: [null, Validators.required],
      Tag: [null, Validators.required],
      // Name: [null, [Validators.required, Validators.minLength(5)]],
      // City: [null, [Validators.required, Validators.minLength(3)]],
    })
  }

  getTagById(id: number) {
    this.apiService.getTagById(id)
    .subscribe((data:any)=>{
      console.log(data);
      this.addTagForm.patchValue({
        name: data.Name,
        description: data.description,
        SlackUrl: data.slack_url,
        Tag: data.tag,
      })
      // this.tagsData = data;
      // this.articles = data['articles'];
    });
  }


  onSubmitForm() {
    
    const formvalue =  this.addTagForm.getRawValue()
    console.log(formvalue);

    let payload = new Tags();

    payload.name = formvalue.name; 
    payload.description = formvalue.description; 
    payload.slack_url = formvalue.SlackUrl; 
    payload.tag = formvalue.Tag; 

    this.apiService.saveTags(payload)
    .subscribe((response: any) => {
      console.log(response);
      console.log(response.status,'response.status');
      if (response.status == 200 ) {
        this.toastr.success('Data Saved Successfully');
        this.addTagForm.reset()
      } else {
        this.toastr.error('Some Error Occured');
      }
    },err => {
        //handle errors here
        this.toastr.error(err);
    });
  }

  onUpdateForm() {
    
    const formvalue =  this.addTagForm.getRawValue()
    console.log(formvalue);

    let payload = new Tags();

    payload.name = formvalue.name; 
    payload.description = formvalue.description; 
    payload.slack_url = formvalue.SlackUrl; 
    payload.tag = formvalue.Tag; 

    this.apiService.updateTags(payload, this.tagID)
    .subscribe((response: any) => {
      console.log(response);
      if (response.status == 200 ) {
        this.toastr.success('Data Saved Successfully');
        this.addTagForm.reset()
      } else {
        this.toastr.error('Some Error Occured');
      }
    },err => {
      //handle errors here
      this.toastr.error(err);
    });
  }
}

export class Tags {
  name: string;
  description: string;
  slack_url: string;
  tag: string;
}
