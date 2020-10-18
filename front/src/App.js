import React, { Component } from 'react';

import './App.css';

import Host from './components/host';
import api from './api';
import 'semantic-ui-css/semantic.min.css'
import { Card, Segment, Divider, Dropdown, Sticky, Menu } from 'semantic-ui-react'

import Shuffle from 'shufflejs'

const vmTypes = [
{key: "all", text: "--All--", value: ""},
{key: "redis", text: "redis", value: "redis"},
{key: "ivr", text: "ivr", value: "ivr"},
{key: "webservice", text: "webservice", value: "webservice"},
{key: "mongo", text: "mongo", value: "mongo"},
{key: "elk", text: "elk", value: "elk"},
{key: "pgbouncer", text: "pgbouncer", value: "pgbouncer"},
{key: "ldap", text: "ldap", value: "ldap"},
];

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      listHost: [],
      vmTypes: vmTypes,
      vmType: null,
    };
    this.sizer = React.createRef();
    this.element = React.createRef();
    this.contextRef = React.createRef();
  }

  async componentDidMount() {
    var listHost = await api.loadHosts();
    this.setState({
      listHost
    });

    this.shuffle = new Shuffle(this.element.current, {
      itemSelector: '.host-item',
      sizer: this.sizer.current,
    });
  }

  async componentDidUpdate() {
    if (this.shuffle) {
      this.shuffle.resetItems();
    }
  }

  render() {
    const { listHost } = this.state;
    return (
      <div  ref={this.contextRef} className="App container">
      <Sticky context={this.contextRef} style={{backgroundColor: "white"}}>
      <Menu
            attached='top'
            tabular
            style={{ backgroundColor: '#fff', paddingBottom: '1em' }}
          >
      <Dropdown placeholder='Vm type' selection options={this.state.vmTypes} />
      </Menu>
      </Sticky>
        <div ref={this.element}>
        { listHost.map((item)  => <Host data={item} key={item.ID} /> ) }
        </div>
      </div>
    );
  }
}
//        <div ref={this.sizer} className="col-1@xs col-1@sm host-grid__sizer" />

export default App;
