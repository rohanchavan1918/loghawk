import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
// import { Tags } from 'src/classFiles/tags';

@Injectable({
  providedIn: 'root'
})
export class ApiServicesService {

  baseURL: string = 'http://127.0.0.1:8080/';

  constructor(private httpClient: HttpClient) { }

  public getAllTags(){
    return this.httpClient.get(this.baseURL + 'tags');
  }

  public getTagById(Id: number){
    return this.httpClient.get(this.baseURL + 'tags/' + Id);
  }

  public saveTags(tags: Tags): Observable<any> {
    const url = this.baseURL + 'tags';
    return this.httpClient.post<any>(url, tags,{ observe: 'response' });
  }

  public updateTags(tags: Tags, id: number): Observable<any> {
    const url = this.baseURL + 'tags/'+ id;
    return this.httpClient.put<any>(url, tags,{ observe: 'response' });
  }

  public deleteTagById( id: number): Observable<any> {
    const url = this.baseURL + 'tags/'+ id;
    return this.httpClient.delete<any>(url, { observe: 'response' });
  }

  public addTagRules(tags: tagRules): Observable<any> {
    const url = this.baseURL + 'tag-rules';
    return this.httpClient.post<any>(url, tags,{ observe: 'response' });
  }
  
  public getStats(): Observable<any> {
    return this.httpClient.get(this.baseURL+ 'stats');
  }
   
  public getLogs(): Observable<any> {
    return this.httpClient.get(this.baseURL+ 'logs');
  }
}

export class Tags {
  name: string;
  description: string;
  slack_url: string;
  tag: string;
}

export class tagRules {
  match_type: string;
  match_value: string;
  priority: number;
  tag_id: number;
}