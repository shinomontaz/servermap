import React, { Component } from "react";
import { Header, Popup, Table } from 'semantic-ui-react'

class Vm extends Component {
  render() {
      var { data } = this.props;
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
