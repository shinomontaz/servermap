import React, { Component } from 'react';

import './App.css';

import Host from './components/host';
import api from './api';
import 'semantic-ui-css/semantic.min.css'

import { Card } from 'semantic-ui-react'

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      listHost: []
    };
  }

  async componentDidMount() {
    var listHost = await api.loadHosts();
    this.setState({
      listHost
    });
  }

  render() {
    const { listHost } = this.state;
    return (
      <div>
      <Card.Group>
      { listHost.map((item)  => <Host data={item} key={item.ID} /> ) }
      </Card.Group>
      </div>
    );
  }
}

export default App;
