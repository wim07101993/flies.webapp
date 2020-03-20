import { NgModule } from "@angular/core";
import { Routes, RouterModule } from "@angular/router";
import { ParticipantsComponent } from "./participants/participants.component";

const routes: Routes = [
  { path: ":year", component: ParticipantsComponent },
  { path: "**", component: ParticipantsComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {}
