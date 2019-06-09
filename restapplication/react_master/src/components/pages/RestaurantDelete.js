import React, { Component } from 'react';
import { HashRouter as Router, Route, NavLink } from 'react-router-dom';

import './UpdationAndDeletion.css';
class ResturantDelete extends Component {
    render() {
      return (
        <Router basename="/restaurant/">
        <div className="App">
            {/* <div className="App__Aside"></div>
            <div className="App__Form">
            <div className="PageSwitcher">
                  <NavLink exact to="/" activeClassName="PageSwitcher__Item--Active" className="PageSwitcher__Item">Resataurants</NavLink>
                  <NavLink exact to="/registration" activeClassName="PageSwitcher__Item--Active" className="PageSwitcher__Item">New Restaurant</NavLink>
                  <NavLink exact to="/UpdateAndDelete" activeClassName="PageSwitcher__Item--Active" className="PageSwitcher__Item">Update Restaurant</NavLink>
                  </div>
           */}
        <div>
          {/* Header Component */}
          {/* <Header /> */}
          {/* Table Component */}
  
          <div className="FormTitle">
                    <NavLink to="/update" activeClassName="FormTitle__Link--Active" className="FormTitle__Link">Update</NavLink> or <NavLink exact to="/delete" activeClassName="FormTitle__Link--Active" className="FormTitle__Link">Delete</NavLink>
                </div> 
          {/* <Route exact path="/" component={Table}>
                </Route>
          <Route exact path="/registration" component={restaurantRegistration}>
                </Route>
                      */}
        </div>
        </div>
        {/* </div> */}
        </Router>
      );
    }
  }
  
  export default ResturantDelete;
  