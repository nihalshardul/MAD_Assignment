import React, { Component } from 'react';
import { Wrapper, AgroTable, Th } from "./table.style";
import faker from 'faker';
import axios from "axios";
//import { Input } from "../input/input.style"

class Table extends Component {
  initialState = {
    searchTerm: '',
    restaurantList:[]
  };

  state = this.initialState;
  componentDidMount() {
    axios.get("http://localhost:8080/restaurant/").then(response => {
      this.setState({restaurantList : response.data.responseData.restaurants})
    })
  }
  // onSearch = ev => {
  //   let restaurantList;
  //   const searchTerm = ev.target.value;
  //   if (!searchTerm) {
  //     // restaurantList = listObj;
  //   axios.get("http://localhost:8080/restaurant/").then(response => {
  //     restaurantList = response.data.responseData.restaurants
  //   })
  //   } else {
  //     restaurantList = this.state.restaurantList.filter(item => {
  //       return item.name.toLowerCase().includes(searchTerm.toLowerCase());
  //     })
  //   }
  //   this.setState({ searchTerm, restaurantList });
  // }

  renderList = lists => {
    return lists.map(list => {
      return (
        <tr>
          {/* <td>{list.id}</td> */}
          <td>{list.name}</td>
          <td>{list.type_of_food}</td>
          <td>{list.address}</td>
          <td>{list.addressline2}</td>
          <td>{list.outcode}</td>
          <td>{list.postcode}</td>
          <td>{list.rating}</td>
      
          <td>{
          <div class="ui card">
            <div class="image">
              <a href={list.url} class="ui medium image">
                <img src={faker.image.avatar()} alt="food"/>
              </a>
            </div>
            <div class="content">
              <div class="description">
                {list.name}
              </div>
            </div>
          </div>
        }
        </td>
        </tr>
      )
    })
  }

  render() {
    return (
      <>
        {/* <Input type="text" placeholder="search by name" onChange={this.onSearch} /> */}
        <Wrapper>
          <AgroTable>
            <thead>
              <tr>
                {/* <Th>Id</Th> */}
                <Th>Name</Th>
                <Th>Type Of Food</Th>
                <Th>Address</Th>
                <Th>AddressLine2</Th>
                <Th>Outcode</Th>
                <Th>Postcode</Th>
                <Th>Rating</Th>
                <Th>Card</Th>
              </tr>
            </thead>
            <tbody>
              {this.renderList(this.state.restaurantList)}
            </tbody>
          </AgroTable>
        </Wrapper>
      </>
    );
  }
}

export default Table;
