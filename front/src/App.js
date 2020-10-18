import React, { Component } from 'react';

import './App.css';

import Host from './components/host';
import api from './api';
import 'semantic-ui-css/semantic.min.css'
import { Dropdown, Sticky, Menu } from 'semantic-ui-react'
import Masonry from 'react-masonry-component';

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
{key: "rabbit", text: "rabbit", value: "rabbit"}
];

const masonryOptions = {
  itemSelector: ".host-item",
  columnWidth: 220,
  fitWidth: true
};

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      listHost: [],
      vmTypes: vmTypes,
      vmType: vmTypes[0].value,
      maxVms: 0,
      meanVms: 0
    };
    this.sizer = React.createRef();
    this.element = React.createRef();
    this.contextRef = React.createRef();
  }

  async componentDidMount() {
    var listHost = await api.loadHosts();

    listHost.sort((h1, h2) => {
      return h2.Vms.length - h1.Vms.length;
    });

    this.state.maxVms = listHost[1].Vms.length;

    listHost.sort(() => { return 0.5 - Math.random() });

    // this.state.maxVms = listHost.reduce((acc,host)=>{
    //   return acc > host.Vms.length ? acc : host.Vms.length
    // },0);

    this.setState({
      listHost
    });

    // this.shuffle = new Shuffle(this.element.current, {
    //   itemSelector: '.host-item',
    //   sizer: this.sizer.current,
    // });
  }

  async componentDidUpdate() {
    // if (this.shuffle) {
    //   this.shuffle.resetItems();
    // }
  }

  handleChangeVmType = (e, { value }) => this.setState({ vmType: value });

  render() {
    const { listHost, vmType, maxVms, meanVms } = this.state;
    return (
      <div ref={this.contextRef} className="App">
      <Sticky context={this.contextRef} style={{backgroundColor: "white"}}>
      <Menu
            attached='top'
            tabular
            style={{ backgroundColor: '#fff', paddingBottom: '1em' }}
          >
      <Dropdown selection options={this.state.vmTypes} value={vmType} onChange={this.handleChangeVmType}/>
      </Menu>
      </Sticky>
        <div ref={this.element}>
        <Masonry className="grid" options={masonryOptions}>
          { listHost.map((item)  => <Host data={item} key={item.ID} types={this.state.vmTypes} maxVms={maxVms} meanVms={meanVms} /> ) }
        </Masonry>
        </div>
      </div>
    );
  }
}

export default App;
