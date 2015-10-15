var Header=React.createClass({

 render: function() {
        return (
          <h1>Header</h1>
        );
    }
});
var Footer=React.createClass({

 render: function() {
        return (
          <h1>Footer</h1>
        );
    }
});
var SearchBar = React.createClass({
    render: function() {
        return (
            <form >
                <input className="serach" type="text"/>
            </form>
        );
    }
});

var FilterableProductTable = React.createClass({
    render: function() {
        return (
            <div>
            <Header/>
                <SearchBar />
            <Footer/>
            </div>
        );
    }
});

ReactDOM.render(
    <FilterableProductTable  />,
    document.getElementById('container')
);