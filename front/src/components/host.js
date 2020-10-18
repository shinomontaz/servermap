import React, { Component } from "react";
import { Header, Segment } from 'semantic-ui-react'
import Vm from './vm';
import Shuffle from 'shufflejs'

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
    this.shuffle = new Shuffle(this.element.current, {
      itemSelector: '.vm-item',
      sizer: this.sizer.current,
    });
  }

  async componentDidUpdate() {
    if (this.shuffle) {
      this.shuffle.resetItems();
    }
  }

  render() {
    const { data } = this.props;
    return (
      <Segment className="host-item" floated='left' style={{minWidth: "15%", width: data.Vms.length > 12 ? "40%" : "20%" }}>
        <Header as='h5'>{data.Name}</Header>
        <div className='meta'>
        {data.Address} ({data.Comment})
        </div>
        <div>{data.Cpu.Cores}x Core {data.Cpu.Name}</div>
        <div>{data.Memory}</div>
        <div ref={this.element}>
        { data.Vms.map((item)  => <Vm data={item} key={item.ID} /> ) }
        </div>
      </Segment>
    );
  }
}

//      <Segment className="col-3@xs col-4@sm host-item" compact style={{maxWidth: "30%"}}>

export default Host;
