import { Component, OnInit, ViewChild } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatSort } from '@angular/material';
import { ParticipantDTO } from './participant';

@Component({
  selector: 'app-participants',
  templateUrl: './participants.component.html',
  styleUrls: ['./participants.component.css']
})
export class ParticipantsComponent implements OnInit {

  displayedColumns: string[] = ['name', 'score'];
  dataSource = new MatTableDataSource(PARTICIPANT_DATA);

  @ViewChild(MatSort, {static: true}) sort: MatSort;

  constructor() { }

  ngOnInit() {
    this.dataSource.sort = this.sort;
  }

  applyFilter(filterValue: string) {
    this.dataSource.filter = filterValue.trim().toLowerCase();
  }

  /** Gets the total amount of flies squashed. */
  getTotalSquashedFlies() {
    return PARTICIPANT_DATA
      .map(t => t.score)
      .reduce((acc, value) => acc + value, 0);
  }

  increase(participant: ParticipantDTO) {
    participant.score++;
  }
  decrease(participant: ParticipantDTO) {
    participant.score--;
  }
}

const PARTICIPANT_DATA: ParticipantDTO[] = [
  {name: "Wim", score: 5},
  {name: "Dimitri", score: 38},
  {name: "Jan G.", score: 20},
  {name: "Frederick", score: 42},
]