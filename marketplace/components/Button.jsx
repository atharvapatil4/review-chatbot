function Button({ onClick, children }) {
    return (
      <button onClick={onClick} className="black_btn">
        {children}
      </button>
    );
  }
  
  export default Button;