import { Injectable } from "@angular/core";
import { HttpClient } from "@angular/common/http";
import { Observable } from "rxjs";
import { ParticipantDTO } from "./participant";

@Injectable({
  providedIn: "root"
})
export class ParticipantsService {
  private apiUrl: string = "http://10.101.90.59:5000/api";

  constructor(private http: HttpClient) {}

  createParticipant(participant: ParticipantDTO) {
    return this.http.post(`${this.apiUrl}/participants`, participant);
  }

  getParticipants(): Observable<ParticipantDTO[]> {
    return this.http.get<ParticipantDTO[]>(`${this.apiUrl}/participants`);
  }

  getParticipant(id: number): Observable<ParticipantDTO> {
    return this.http.get<ParticipantDTO>(`${this.apiUrl}/participants/${id}`);
  }

  updateName(id: number, newName: string): Observable<ParticipantDTO> {
    return this.http.put<ParticipantDTO>(
      `${this.apiUrl}/participants/${id}/name/?name=${newName}`,
      ""
    );
  }

  updateScore(id: number, score: number): Observable<ParticipantDTO> {
    return this.http.put<ParticipantDTO>(
      `${this.apiUrl}/participants/${id}/score/?score=${score}`,
      ""
    );
  }

  deleteParticipant(id: number) {
    return this.http.delete<ParticipantDTO>(
      `${this.apiUrl}/participants/${id}`
    );
  }
}
