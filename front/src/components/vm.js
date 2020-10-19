import React, { Component } from "react";
import { Header, Popup, Segment, Label, Icon, Table } from 'semantic-ui-react'
import Ram from './ram';
import Cpu from './cpu';

class Vm extends Component {
  render() {
      var { data, types, maxRam, maxCpu } = this.props;
      return (
      <Table.Row>
      <Table.Cell>
      <Popup
        trigger={<Header as='h5'>{data.Name}</Header>}
        content={data.OperatingSystem.Distribution}
        basic
      />
      </Table.Cell>
      <Table.Cell>
      {data.Memory}
      </Table.Cell>
      <Table.Cell>
      {data.Cpu.Cores}
      </Table.Cell>
      </Table.Row>
      );
  }
}

export default Vm;
