import { Component, OnInit, ViewChild, Inject } from "@angular/core";
import { MatTableDataSource } from "@angular/material/table";
import { MatSort } from "@angular/material";
import { ParticipantDTO } from "./participant";
import { ParticipantsService } from "./participants.service";
import {
  MatDialog,
  MatDialogRef,
  MAT_DIALOG_DATA
} from "@angular/material/dialog";
import { ActivatedRoute } from "@angular/router";

@Component({
  selector: "app-participants",
  templateUrl: "./participants.component.html",
  styleUrls: ["./participants.component.scss"]
})
export class ParticipantsComponent implements OnInit {
  displayedColumns: string[] = [
    //'id',
    "name",
    "score",
    "buttons"
  ];
  dataSource = new MatTableDataSource([]);
  private year: number;

  @ViewChild(MatSort, { static: true }) sort: MatSort;

  constructor(
    private service: ParticipantsService,
    public dialog: MatDialog,
    private route: ActivatedRoute
  ) {}

  ngOnInit() {
    this.route.params.subscribe(params => {
      this.year = +params["year"]; // '+' converts the string to a number
      this.refresh();
    });
  }

  applyFilter(filterValue: string) {
    this.dataSource.filter = filterValue.trim().toLowerCase();
  }

  newParticipant() {
    const dialogRef = this.dialog.open(NewParticipantDialog, {
      width: "250px",
      data: { name: "", score: 0 }
    });

    dialogRef.afterClosed().subscribe(result => {
      if (result == null || result == "") {
        return;
      }

      this.service
        .createParticipant(this.year, { id: 0, name: result, score: 0 })
        .subscribe(p => {
          this.refresh();
        });
    });
  }

  getTotalSquashedFlies() {
    return this.dataSource.data
      .map(t => t.score)
      .reduce((acc, value) => acc + value, 0);
  }

  increaseParticipantScore(participant: ParticipantDTO) {
    this.updateParticipantScore(participant, participant.score + 1);
  }
  decreaseParticipantScore(participant: ParticipantDTO) {
    this.updateParticipantScore(participant, participant.score - 1);
  }

  updateParticipantScore(participant: ParticipantDTO, score: number) {
    this.service.updateScore(this.year, participant.id, score).subscribe(p => {
      participant.score = p.score;
      participant.name = p.name;
    });
  }

  refresh() {
    this.service.getParticipants(this.year).subscribe(ps => {
      this.dataSource = new MatTableDataSource(ps);
      this.dataSource.sort = this.sort;
    });
  }
}

@Component({
  selector: "app-new-participant-dialog",
  templateUrl: "./new-participant.dialog.html"
})
export class NewParticipantDialog {
  constructor(
    public dialogRef: MatDialogRef<NewParticipantDialog>,
    @Inject(MAT_DIALOG_DATA) public data: ParticipantDTO
  ) {}

  onCancelClick(): void {
    this.dialogRef.close();
  }
}
