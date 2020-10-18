import React, { Component } from "react";
import { Card, Segment, List } from 'semantic-ui-react'
import Vm from './vm';

class Host extends Component {
  render() {
    const { data } = this.props;
    return (
      <Card>
            <Card.Content>
            <Card.Header>{data.Name}</Card.Header>
            <Card.Meta>
              <span className='address'>{data.Address}</span>
              <span className='comment'>({data.Comment})</span>
            </Card.Meta>
            <Card.Description>
            <div>{data.Cpu.Cores}xCore {data.Cpu.Name}</div>
            <div>{data.Memory}</div>
            </Card.Description>
            <Card.Content extra>
            <Card.Group itemsPerRow={2}>
            { data.Vms.map((item)  => <Vm data={item} key={item.ID} /> ) }
            </Card.Group>
            </Card.Content>
          </Card.Content>
          </Card>
    );
  }
}

//            <Card.Group itemsPerRow={ data.Vms.length == 0 ? 1 : data.Vms.length > 12 ? 4 : data.Vms.length > 8 ? 3 : 2}>

export default Host;
