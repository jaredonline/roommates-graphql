import React from 'react';
import Relay from 'react-relay';

class PersonBalance extends React.Component {
  render() {
    return (
        <div>
            <div className='balance'>
                <div className='span2'>
                    <span className='balance-name'>{this.props.person.name}</span>
                </div>
                <div className='span1 offset1'>
                    <span className='balance-amount'> $928.00 </span>
                </div>
            </div>
            <div className="clearfix"></div>
        </div>
    );
  }
}

export default Relay.createContainer(PersonBalance, {
  fragments: {
    person: () => Relay.QL`
        fragment on Person {
            name
        }
    `,
  },
});
