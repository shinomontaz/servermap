import React, { Component } from 'react';

import './App.css';

import Host from './components/host';
import api from './api';
import 'semantic-ui-css/semantic.min.css'
import { Dropdown, Sticky, Menu } from 'semantic-ui-react'
import Masonry from 'react-masonry-component';

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
};

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      listDC: [],
      vmTypes: vmTypes,
      vmType: vmTypes[0].value,
      maxVms: 0,
      maxCpu: 0,
      maxRam: 0,
      openAll: false,
      hidden: true,
    };
    this.element = React.createRef();
    this.contextRef = React.createRef();
  }

  prepareListDc = ( listHost ) => {
    return listHost.reduce(function(r, h){
    var k = h.Comment.trim();
    if (r[k] || (r[k]=[])) r[k].push(h);
    return r;
  }, {});
  }

  async componentDidMount() {
    var listHost = await api.loadHosts(this.state.vmType);

    listHost.sort((h1, h2) => {
      return h2.Vms.length - h1.Vms.length;
    });

    this.state.maxVms = listHost[1].Vms.length;

    listHost.sort(() => { return 0.5 - Math.random() });

    this.state.maxCpu = listHost.reduce((acc,host)=>{
      return acc > host.Cpu.Cores * host.Cpu.Threads ? acc : host.Cpu.Cores * host.Cpu.Threads
    },0);

    this.state.maxRam = listHost.reduce((acc,host)=>{
      return acc > host.RawMem ? acc : host.RawMem
    },0);

    this.setState({
      listDC: this.prepareListDc(listHost)
    });
  }

  async componentDidUpdate(prevProps, prevState) {
    if (this.state.vmType !== prevState.vmType) {
      var listHost = await api.loadHosts(this.state.vmType);

      this.setState({
        listDC: this.prepareListDc(listHost),
        openAll: this.state.vmType !== "" ? true : false
      });
    }
  }

  handleChangeVmType = (e, { value }) => this.setState({ vmType: value });

  handleContentsLoaded = () => this.setState({ hidden: false });

  render() {
    const { listDC, vmType, openAll, hidden, maxRam, maxCpu } = this.state;
    return (
      <div ref={this.contextRef} className="App">
      <Sticky context={this.contextRef} style={{backgroundColor: "white"}}>
      <Menu
          attached='top'
          tabular
          style={{ backgroundColor: '#fff', paddingBottom: '1em' }}
      >
      <Dropdown selection options={this.state.vmTypes} value={vmType} onChange={this.handleChangeVmType}  style={{backgroundColor: "white"}}/>
      </Menu>
      </Sticky>
        <div ref={this.element}>
        {
          Object.entries(listDC).map( ([key, dc]) => (
            <div>
            <h1>{key}</h1>
            <Masonry
              className="grid"
              options={masonryOptions}
              onImagesLoaded={this.handleContentsLoaded}
              visibility = {hidden}
            >
            { dc.map((item) => <Host data={item} key={item.ID} types={this.state.vmTypes} maxRam={maxRam} maxCpu={maxCpu} open={openAll} /> ) }
            </Masonry>
            <hr/>
            </div>
          ))
        }
        </div>
      </div>
    );
  }
}

export default App;
