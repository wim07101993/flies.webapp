import { Component, OnInit } from "@angular/core";

@Component({
  selector: "app-navigation",
  templateUrl: "./navigation.component.html",
  styleUrls: ["./navigation.component.scss"]
})
export class NavigationComponent implements OnInit {
  startYear = 2019;
  years: number[] = [];

  constructor() {
    var currentYear = new Date().getFullYear();
    for (let year = this.startYear; year <= currentYear; year++) {
      this.years.push(year);
    }
  }

  ngOnInit() {}
}
