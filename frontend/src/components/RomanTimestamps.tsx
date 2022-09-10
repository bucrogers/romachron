import React from "react";
import axios from "axios";

function RomanTimestamps() {
    const [posts, setPosts] = React.useState<any>([]);

    React.useEffect(() => {
        const interval = setInterval(() => {
            axios
              .get("http://localhost:8080/romadatetime")
              .then((response) => setPosts(response.data));
        }, 500);
        return () => clearInterval(interval);
    }, []);

    return (
        <ul className="posts">
          <table align="center">
            <tr>
              <td width="200" align="left" valign="top">
                <h2>{posts.hr_12h_format + ":" + posts.min +  ":" + posts.sec}</h2>
              </td>
              <td align="left" valign="top">
              <h2>{posts.am_pm}</h2>
              </td>
            </tr>
          </table>  
          <h2>{posts.day_expression + ", " + posts.date_expression}</h2>
        </ul>
    );
}

export default RomanTimestamps;