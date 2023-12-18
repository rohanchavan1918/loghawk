import { Component, OnInit } from '@angular/core';
import { ApiServicesService } from '../services/api-services.service';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ToastrService } from 'ngx-toastr';
// import { MatchType } from 'src/enum/matchType.enum';

@Component({
  selector: 'app-table-list',
  templateUrl: './table-list.component.html',
  styleUrls: ['./table-list.component.css']
})
export class TableListComponent implements OnInit {

  tagsData: any[] = [];
  tagRules: any[] = [];
  addTagRuleForm: FormGroup;
  tagId: number;
  showForm: boolean = false;
  deleteId: number

  public TestEnum = MatchType;

  constructor(
    private apiService: ApiServicesService,
    private fb: FormBuilder,
    private toastr: ToastrService
  ) { }


  ngOnInit() {
    this.initializeForm();
    this.getAllTags()
  }

  getAllTags() {
    this.apiService.getAllTags().subscribe((data:any)=>{
      console.log(data);
      this.tagsData = data;
    });
  }

  initializeForm() {
    this.addTagRuleForm = this.fb.group({
      matchType: [null, Validators.required],
      matchValue: [null, Validators.required],
      priority: [null, Validators.required]
    })
  }

  submitMatchTypeForm() {
    
    const formvalue =  this.addTagRuleForm.getRawValue()
    console.log(formvalue);

    let payload = new tagRules();

    payload.match_type = formvalue.matchType; 
    payload.match_value = formvalue.matchValue; 
    payload.priority = +formvalue.priority; 
    payload.tag_id = this.tagId; 

    this.apiService.addTagRules(payload)
    .subscribe((response: any) => {
      console.log(response);
      if (response.status == 200 ) {
        this.toastr.success('Data Saved Successfully');
        this.addTagRuleForm.reset()
      } else {
        this.toastr.error('Some Error Occured');
      }
    },err => {
      //handle errors here
      this.toastr.error(err);
    });
  }

  deleteTag(id) {

    this.apiService.deleteTagById(id)
    .subscribe((response: any) => {
      console.log(response);
      if (response.status == 200 ) {
        this.toastr.success('Data Deleted Successfully');
        this.getAllTags()
      } else {
        this.toastr.error('Some Error Occured');
      }
    },err => {
      //handle errors here
      this.toastr.error(err);
    });

  }

  copyText(val: string){
    let selBox = document.createElement('textarea');
      selBox.style.position = 'fixed';
      selBox.style.left = '0';
      selBox.style.top = '0';
      selBox.style.opacity = '0';
      selBox.value = val;
      document.body.appendChild(selBox);
      selBox.focus();
      selBox.select();
      document.execCommand('copy');
      document.body.removeChild(selBox);
    }

}

export enum MatchType {
	contains = "Contains",
	starts_with = "Starts With",
	ends_with = "Ends With",
	regex = "Regex",
  custom = "Custom"
}



export class tagRules {
  match_type: string;
  match_value: string;
  priority: number;
  tag_id: number;
}