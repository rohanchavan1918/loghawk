import { Component, OnInit } from '@angular/core';
import { ApiServicesService } from '../services/api-services.service';

@Component({
  selector: 'app-Log-List',
  templateUrl: './Log-List.component.html',
  styleUrls: ['./Log-List.component.scss']
})
export class LogListComponent implements OnInit {

  statsData: any[] = null;
  searchText: string = ''
  constructor(
    private apiService: ApiServicesService
  ) { }

  ngOnInit() {
    this.getAllLogs()
  }

  getAllLogs() {
    this.apiService.getLogs()
    .subscribe((data:any)=>{
      console.log(data);
      this.statsData = data;
      // this.articles = data['articles'];
    });
  }
}
