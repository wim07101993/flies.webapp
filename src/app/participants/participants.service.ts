import { Injectable } from "@angular/core";
import { HttpClient } from "@angular/common/http";
import { Observable } from "rxjs";
import { ParticipantDTO } from "./participant";

@Injectable({
  providedIn: "root"
})
export class ParticipantsService {
  private apiUrl: string = "http://localhost:5000/api";

  constructor(private http: HttpClient) {}

  createParticipant(year: number, participant: ParticipantDTO) {
    return this.http.post(`${this.apiUrl}/participants/${year}/`, participant);
  }

  getParticipants(year: number): Observable<ParticipantDTO[]> {
    return this.http.get<ParticipantDTO[]>(
      `${this.apiUrl}/participants/${year}`
    );
  }

  getParticipant(year: number, id: number): Observable<ParticipantDTO> {
    return this.http.get<ParticipantDTO>(
      `${this.apiUrl}/participants/${year}/${id}`
    );
  }

  updateName(
    year: number,
    id: number,
    newName: string
  ): Observable<ParticipantDTO> {
    return this.http.put<ParticipantDTO>(
      `${this.apiUrl}/participants/${year}/${id}/name/?name=${newName}`,
      ""
    );
  }

  updateScore(
    year: number,
    id: number,
    score: number
  ): Observable<ParticipantDTO> {
    return this.http.put<ParticipantDTO>(
      `${this.apiUrl}/participants/${year}/${id}/score/?score=${score}`,
      ""
    );
  }

  deleteParticipant(year: number, id: number) {
    return this.http.delete<ParticipantDTO>(
      `${this.apiUrl}/participants/${year}/${id}`
    );
  }
}
