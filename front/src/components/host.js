import React, { Component } from "react";
import { Header, Segment, Transition, Popup, Table } from 'semantic-ui-react'
import Vm from './vm';
import Ram from './ram';
import Cpu from './cpu';

class Host extends Component {
  constructor(props) {
    super(props);
    this.state = {
      listVm: [],
      open: props.open,
    };

    props.data.Vms.sort((v1, v2) => {
      return v2.Cpu.Cores - v1.Cpu.Cores;
    });
  }

  async componentDidMount() {
  }

  toggleOpen = () => this.setState((prevState) => ({ open: !prevState.open }))

  render() {
    const { data, types, maxRam, maxCpu } = this.props;
    const isSpecial = data.Name === 'nonexist' || data.Vms.length === 0
//    const rows = data.Vms.length/ 2;
    return (
      <Segment
      className="host-item"
      style={{minWidth: "10%", maxWidth: "99%" }}
      inverted={isSpecial ? true : false}
      color={isSpecial ? 'red' : null}
      tertiary={isSpecial}
      onClick={this.toggleOpen}
      circular
      >
        <Header as='h5'>{data.Name}</Header>
        <div className='meta'>
        {data.Address}
        </div>
        <Popup
          trigger={<div><Ram total={data.RawMem} maxRam={maxRam} current={data.Vms.reduce((acc,vm)=>{
            return acc + vm.RawMem
          },0)} /></div>}
          content={data.Memory}
          basic
        />
        <Popup
            trigger={<div><Cpu total={data.Cpu.Cores * data.Cpu.Threads} maxCpu={maxCpu} current={data.Vms.reduce((acc,vm)=>{
              return acc + (vm.Cpu.Cores * vm.Cpu.Threads)
            },0)} /></div>}
            content={data.Cpu.Cores+"x Core "+ data.Cpu.Name}
            basic
        />
        <Transition visible={this.state.open} animation='scale'>
        <div>
          <Table className="vm-item"  basic='very' celled collapsing >
          <Table.Header>
             <Table.Row>
               <Table.HeaderCell></Table.HeaderCell>
               <Table.HeaderCell>Ram</Table.HeaderCell>
               <Table.HeaderCell>vCpu</Table.HeaderCell>
             </Table.Row>
           </Table.Header>
            <Table.Body>
            {data.Vms.map((item) => <Vm data={item} key={item.ID} types={types} maxRam={data.RawMem} maxCpu={data.Cpu.Cores * data.Cpu.Threads} /> )}
            </Table.Body>
          </Table>
        </div>
        </Transition>
      </Segment>
    );
  }
}

export default Host;
