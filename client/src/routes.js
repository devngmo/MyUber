/* 
* Generated by
* 
*      _____ _          __  __      _     _
*     / ____| |        / _|/ _|    | |   | |
*    | (___ | | ____ _| |_| |_ ___ | | __| | ___ _ __
*     \___ \| |/ / _` |  _|  _/ _ \| |/ _` |/ _ \ '__|
*     ____) |   < (_| | | | || (_) | | (_| |  __/ |
*    |_____/|_|\_\__,_|_| |_| \___/|_|\__,_|\___|_|
*
* The code generator that works in many programming languages
*
*			https://www.skaffolder.com
*
*
* You can generate the code from the command-line
*       https://npmjs.com/package/skaffolder-cli
*
*       npm install -g skaffodler-cli
*
*   *   *   *   *   *   *   *   *   *   *   *   *   *   *   *
*
* To remove this comment please upgrade your plan here: 
*      https://app.skaffolder.com/#!/upgrade
*
* Or get up to 70% discount sharing your unique link:
*       https://app.skaffolder.com/#!/register?friend=5d525f5a7b158e50f28eb4af
*
* You will get 10% discount for each one of your friends
* 
*/
// Dependencies
import React, { Component } from "react";
import { Fragment } from "react";
import { Route, Switch } from "react-router";
import { PrivateRoute } from "./security/PrivateRoute";

// Material UI
import Paper from "@material-ui/core/Paper";

/* START MY VIEWS IMPORT */

import BookingEdit from "./pages/BookingEdit";
import BookingList from "./pages/BookingList";
import CarEdit from "./pages/CarEdit";
import CarList from "./pages/CarList";
import CarRegister from "./pages/CarRegister";
import CarStatus from "./pages/CarStatus";
import DriverEdit from "./pages/DriverEdit";
import DriverList from "./pages/DriverList";
import TrackEdit from "./pages/TrackEdit";
import TrackList from "./pages/TrackList";

/* END MY VIEWS IMPORT */

// CUSTOM VIEWS
import Login from "./pages/Login";
import Home from "./pages/Home";
import Profile from "./pages/Profile";
import UserEdit from "./pages/UserEdit";
import UserList from "./pages/UserList";

class Routes extends Component {
  render() {
    return (
      <Switch>
        <Fragment>
          <Paper>
            <div className="main-cointainer">
              <Route exact path="/login" component={Login} />
              <PrivateRoute exact path="/" component={Home} />
              <PrivateRoute exact path="/profile" component={Profile} />
              <PrivateRoute exact path="/users/:id" component={UserEdit} roles={["ADMIN"]}/>
              <PrivateRoute exact path="/users" component={UserList} roles={["ADMIN"]}/>
              
              {/* CUSTOM VIEWS */}

              <PrivateRoute exact path="/home" component={Home} />

              {/* START MY VIEWS */}

              <PrivateRoute exact path="/bookings/:id" component={ BookingEdit }  />
              <PrivateRoute exact path="/bookings" component={ BookingList }  />
              <PrivateRoute exact path="/cars/:id" component={ CarEdit }  />
              <PrivateRoute exact path="/cars" component={ CarList }  />
              <PrivateRoute exact path="/car/register" component={ CarRegister }  />
              <PrivateRoute exact path="/car/status/:id" component={ CarStatus }  />
              <PrivateRoute exact path="/drivers/:id" component={ DriverEdit }  />
              <PrivateRoute exact path="/drivers" component={ DriverList }  />
              <PrivateRoute exact path="/tracks/:id" component={ TrackEdit }  />
              <PrivateRoute exact path="/tracks" component={ TrackList }  />

             {/* END MY VIEWS */}

            </div>
          </Paper>
        </Fragment>
      </Switch>
    );
  }
}

export default Routes;