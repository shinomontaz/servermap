import React, { Component } from "react";
import { Menu, Card, Segment, Label, Icon } from 'semantic-ui-react'

class Vm extends Component {
  getDataGroups = (types, name) => {
    var currTypes = types.filter(type => name.includes(type.value)).map(filteredType => filteredType.key);
    return currTypes;
  }

  render() {
      const { data, types } = this.props;
      return (
    <Segment className="vm-item" data-groups={this.getDataGroups(types, data.Name)}>
      <Card.Content>
        <Card.Header>{data.Name}</Card.Header>
        <Card.Meta>{data.OperatingSystem.Distribution}</Card.Meta>
      </Card.Content>
      <Card.Content extra>
      <Menu compact>
    <Menu.Item as='a'>
      <Icon name='th' />
      <Label floating>
        {data.Cpu.Cores}
      </Label>
    </Menu.Item>
    <Menu.Item as='a'>
      <Icon name='microchip' />
      <Label floating>
        {data.Memory}
      </Label>
    </Menu.Item>
  </Menu>
      </Card.Content>
    </Segment>
      );
  }
}

export default Vm;
