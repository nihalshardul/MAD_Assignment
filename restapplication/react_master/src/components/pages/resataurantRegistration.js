import React, { Component } from 'react';
// import { Link } from 'react-router-dom';

class restaurantRegistration extends Component {
    constructor() {
        super();

        this.state = {
          name:'',
          address: '',
          addressLine2: '',
          url: '',
          outcode:'',
          postcode:'',
          rating:'',
          type_of_food:'',
          hasAgreed: false
        };

        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    handleChange(e) {
        let target = e.target;
        let value = target.type === 'checkbox' ? target.checked : target.value;
        let name = target.name;

        this.setState({
          [name]: value
        });
    }

    handleSubmit(e) {
        e.preventDefault();

        console.log('The form was submitted with the following data:');
        console.log(this.state);
    }

    render() {
        return (
        <div className="FormCenter">
            <form onSubmit={this.handleSubmit} className="FormFields">
              <div className="FormField">
                <label className="FormField__Label" htmlFor="name">Restaurant Name</label>
                <input type="text" id="name" className="FormField__Input" placeholder="Enter your full name" name="name" value={this.state.name} onChange={this.handleChange} />
              </div>
              <div className="FormField">
                <label className="FormField__Label" htmlFor="address">Adrress</label>
                <input type="text" id="address" className="FormField__Input" placeholder="Enter your Adrress" name="address" value={this.state.address} onChange={this.handleChange} />
              </div>
              <div className="FormField">
                <label className="FormField__Label" htmlFor="addressLine2">Further Adrress</label>
                <input type="text" id="addressLine2" className="FormField__Input" placeholder="Enter your Further Address" name="addressLine2" value={this.state.addressLine2} onChange={this.handleChange} />
              </div>
              <div className="FormField">
                <label className="FormField__Label" htmlFor="url">URL</label>
                <input type="text" id="url" pattern="https?://.+" className="FormField__Input" placeholder="Enter Restaurant URL" name="url" value={this.state.url} onChange={this.handleChange} />
              </div>
              <div className="FormField">
                <label className="FormField__Label" htmlFor="outcode">Outcode</label>
                <input type="text" id="outcode" className="FormField__Input" placeholder="Enter Restaurant Outcode" name="outcode" value={this.state.outcode} onChange={this.handleChange} />
              </div>
              <div className="FormField">
                <label className="FormField__Label" htmlFor="postcode">Postcode</label>
                <input type="text" id="postcode" className="FormField__Input" placeholder="Enter Restaurant Postcode" name="postcode" value={this.state.postcode} onChange={this.handleChange} />
              </div>
              <div className="FormField">
                <label className="FormField__Label" htmlFor="rating">Rating</label>
                <input type='number' step="0.1" min='0' max='5' id="rating" className="FormField__Input" placeholder="Enter Restaurant Rating" name="rating" value={this.state.rating} onChange={this.handleChange} />
              </div>
              <div className="FormField">
                <label className="FormField__Label" htmlFor="type_of_food">Type of Food</label>
                <input type="text" id="type_of_food" className="FormField__Input" placeholder="Enter Types of Food" name="type_of_food" value={this.state.type_of_food} onChange={this.handleChange} />
              </div>
              
              <div className="FormField">
                <label className="FormField__CheckboxLabel">
                    <input className="FormField__Checkbox" type="checkbox" name="hasAgreed" value={this.state.hasAgreed} onChange={this.handleChange} /> I agree all informations are correct
                </label>
              </div>

              <div className="FormField">
                  <button className="FormField__Button mr-20">SUbmit</button> 
              </div>
            </form>
          </div>
        );
    }
}

export default restaurantRegistration;
