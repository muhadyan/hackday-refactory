import axios from "axios";
import { useState, useEffect } from "react";

export default function Table() {
  // decalre state
  const [student, setStudent] = useState(null);
  // call axios in useEffect
  useEffect(() => {
    axios.get("http://localhost:8080/students").then((res) => {
      setStudent(res.data);
    });
  }, []);

  if (!student) return null;
  // insert axios result to state
  // show the state

  return (
    <div className="container">
      <div class="row justify-content-md-center">
        <h1 class="col col-lg-4">List of Students</h1>
      </div>
      <div className="row justify-content-md-center">
        <div className="col-6">
          <table className="table table-hover">
            <thead>
              <tr>
                <th scope="col">Student ID</th>
                <th scope="col">Full Name</th>
                <th scope="col">Extra Name</th>
              </tr>
            </thead>
            <tbody>
              {student.map((value, index) => {
                return (
                  <tr key={index}>
                    <th scope="row">{value.student_id}</th>
                    <td>{value.full_name}</td>
                    <td>{value.extra_name}</td>
                  </tr>
                );
              })}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  );
}
