import React, { Component } from "react";
import { Header, Segment, Transition, Card } from 'semantic-ui-react'
import Vm from './vm';
import Masonry from 'react-masonry-component';

const masonryOptions = {
  itemSelector: ".vm-item",
};

class Host extends Component {
  constructor(props) {
    super(props);
    this.state = {
      listVm: [],
      open: false,
    };
    this.sizer = React.createRef();
    this.element = React.createRef();
  }

  async componentDidMount() {
  }

  async componentDidUpdate() {
  }


  toggleOpen = () => this.setState((prevState) => ({ open: !prevState.open }))

  render() {
    const { data, types, maxVms } = this.props;
    const isSpecial = data.Name == 'nonexist' || data.Vms.length == 0

//    const width = Math.floor( data.Vms.length > 0 ? (( data.Vms.length / maxVms) * 10 ) : 1 ) * 10;

    return (
      <Segment
      className="host-item"
      style={{minWidth: "10%", maxWidth: "99%" }}
      inverted={isSpecial ? true : false}
      color={isSpecial ? 'red' : 'black'}
      tertiary={isSpecial}
      onClick={this.toggleOpen}
      >
        <Header as='h5'>{data.Name}</Header>
        <div className='meta'>
        {data.Address} ({data.Comment})
        </div>
        <div>{data.Cpu.Cores}x Core {data.Cpu.Name}</div>
        <div>{data.Memory}</div>
        <div ref={this.element}>
        <Transition visible={this.state.open} animation='scale'>
          <Card.Group>
          { data.Vms.map((item)  => <Vm data={item} key={item.ID} types={types}/> ) }
          </Card.Group>
        </Transition>
        </div>
      </Segment>
    );
  }
}
//      <Segment className="host-item" compact style={{maxWidth: "30%"}}>

export default Host;
