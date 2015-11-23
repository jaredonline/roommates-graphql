import React from 'react';
import Relay from 'react-relay';
import Balances from './Balances';

class App extends React.Component {
  render() {
    return (
        <div>
            <Balances key={this.props.me.id} person={this.props.me} />
            <section className="bounded">
                <div className="three-quarter-container bounded no-border dark">
                    <div className='add-new-link'>
                        <a href="/expenses/new">Add Expense</a>
                        <a href="/expenses">All Expenses</a>
                    </div>
                    <h3 className='ribbon-header'>Recent Expenses</h3>
                    <div className='expenses'>
                        <div className='expense'>
                            <div className='expense-date'>
                                Nov 09
                            </div>
                            <div className='expense-amount'>
                                $1.00
                                <div className='expense-share'>
                                    your share is
                                    <span className='currency'>$1.00</span>
                                </div>
                            </div>
                            <div className='expense-description'>
                                Allison
                                <span className='expense-small'>paid for</span>
                                This is a test
                                <span className='expense-notes'>
                                    Seriously Just a Test
                                </span>
                                <span className='expense-small'>
                                    <a href="/expenses/1438" data-confirm="Are you sure you want to permanently delete this expense?" data-method="delete" rel="nofollow">delete</a>
                                </span>
                            </div>
                            <div className='clearfix'></div>
                        </div>
                        <div className='clearfix'></div>
                    </div>
                </div>
            </section>
        </div>
    );
  }
}

export default Relay.createContainer(App, {
  fragments: {
    me: () => Relay.QL`
        fragment on Person {
            ${Balances.getFragment('person')}
        }
    `,
  },
});
