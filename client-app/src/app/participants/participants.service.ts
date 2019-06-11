import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { ParticipantDTO } from './participant';

@Injectable({
  providedIn: 'root'
})
export class ParticipantsService {

  private apiUrl: string = 'http://localhost:5000/api'
  
  constructor(
    private http: HttpClient
  ) { }

  createParticipant(participant: ParticipantDTO) {
    return this.http.post(`${this.apiUrl}/participants`, participant);
  }

  getParticipants() : Observable<ParticipantDTO[]> {
    return this.http.get<ParticipantDTO[]>(`${this.apiUrl}/participants`);
  }

  getParticipant(name: string) : Observable<ParticipantDTO> {
    return this.http.get<ParticipantDTO>(`${this.apiUrl}/participants/${name}`)
  }

  updateName(oldName: string, newName: string) : Observable<ParticipantDTO> {
    return this.http.put<ParticipantDTO>(`${this.apiUrl}/participants/${oldName}/name/?name=${newName}`, "");
  }

  updateScore(name: string, score: number) : Observable<ParticipantDTO> {
    return this.http.put<ParticipantDTO>(`${this.apiUrl}/participants/${name}/score/?score=${score}`, "");
  }

  deleteParticipant(name: string) {
    return this.http.delete<ParticipantDTO>(`${this.apiUrl}/participants/${name}`)
  }
}
