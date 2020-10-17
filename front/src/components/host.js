import React, { Component } from "react";
import { Card, List } from 'semantic-ui-react'

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
    </Card.Content>
      </Card>
    );
  }
}

export default Host;
