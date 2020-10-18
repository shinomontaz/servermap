import React, { Component } from "react";
import { Header, Segment } from 'semantic-ui-react'
import Vm from './vm';
import Shuffle from 'shufflejs'
import Masonry from 'react-masonry-component';

const masonryOptions = {
  itemSelector: ".vm-item",
};

class Host extends Component {
  constructor(props) {
    super(props);
    this.state = {
      listVm: []
    };
    this.sizer = React.createRef();
    this.element = React.createRef();
  }

  async componentDidMount() {
    // this.shuffle = new Shuffle(this.element.current, {
    //   itemSelector: '.vm-item',
    //   sizer: this.sizer.current,
    // });
  }

  async componentDidUpdate() {
    if (this.shuffle) {
      this.shuffle.resetItems();
    }
  }

  render() {
    const { data, types, maxVms } = this.props;
    const isSpecial = data.Name == 'nonexist' || data.Vms.length == 0
    console.log( data.Vms.length, maxVms, data.Vms.length > 0 ? (( data.Vms.length / maxVms) * 100 ) + "%" : "10%");

    const width = Math.floor( data.Vms.length > 0 ? (( data.Vms.length / maxVms) * 10 ) : 1 ) * 10;

    return (
      <Segment className="host-item" style={{minWidth: "10%", maxWidth: "99%", width: width + "%" }} inverted={isSpecial ? true : false} color={isSpecial ? 'red' : 'black'} tertiary={isSpecial}>
        <Header as='h5'>{data.Name}</Header>
        <div className='meta'>
        {data.Address} ({data.Comment})
        </div>
        <div>{data.Cpu.Cores}x Core {data.Cpu.Name}</div>
        <div>{data.Memory}</div>
        <div ref={this.element}>
        <Masonry className="grid" options={masonryOptions}>
          { data.Vms.map((item)  => <Vm data={item} key={item.ID} types={types}/> ) }
        </Masonry>
        </div>
      </Segment>
    );
  }
}
//      <Segment className="host-item" compact style={{maxWidth: "30%"}}>

export default Host;
