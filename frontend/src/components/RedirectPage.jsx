import React, { useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { useParams, useLocation } from "react-router-dom";
import { Typography, Box, Container, Paper } from "@mui/material";
import { getOriginalUrl } from "../redux/actions";

const RedirectPage = () => {
  const { code } = useParams();
  const location = useLocation();
  const dispatch = useDispatch();
  const [showCode, setShowCode] = useState(false);
  const { originalUrl, loading, error } = useSelector((state) => state.url);

  useEffect(() => {
    const searchParams = window.location.href;
    const checkOnly = searchParams[searchParams.length - 1] === "_";
    setShowCode(checkOnly);
  }, []);

  useEffect(() => {
    const searchParams = window.location.href;
    const checkOnly = searchParams[searchParams.length - 1] === "_";
    console.log(searchParams);
    dispatch(getOriginalUrl(code, checkOnly));
  }, [dispatch, code, location]);

  useEffect(() => {
    const searchParams = window.location.href;
    const checkOnly = searchParams[searchParams.length - 1] === "_";
    if (originalUrl && !checkOnly) {
      window.location.href = originalUrl;
    }
  }, [originalUrl, location]);

  if (loading) {
    return (
      <Container
        maxWidth="sm"
        sx={{
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
          minHeight: "100vh",
          position: "relative",
          zIndex: 1,
        }}
      ></Container>
    );
  }

  if (error) {
    return (
      <Container
        maxWidth="sm"
        sx={{
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
          minHeight: "100vh",
          position: "relative",
          zIndex: 1,
        }}
      >
        <Paper
          elevation={3}
          sx={{
            p: 4,
            backgroundColor: "background.paper",
            border: "2px solid",
            borderColor: "primary.main",
          }}
        >
          <Typography color="error" align="center">
            {error}
          </Typography>
        </Paper>
      </Container>
    );
  }

  if (originalUrl && showCode) {
    return (
      <Container
        maxWidth="sm"
        sx={{
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
          minHeight: "100vh",
          position: "relative",
          zIndex: 1,
        }}
      >
        <Paper
          elevation={3}
          sx={{
            p: 4,
            backgroundColor: "background.paper",
            border: "2px solid",
            borderColor: "primary.main",
          }}
        >
          <Typography
            variant="h6"
            align="center"
            sx={{ color: "text.primary" }}
          >
            This short URL redirects to:
          </Typography>
          <Typography
            variant="body1"
            align="center"
            sx={{ mt: 2, color: "primary.main", wordBreak: "break-all" }}
          >
            {originalUrl}
          </Typography>
        </Paper>
      </Container>
    );
  }

  return null;
};

export default RedirectPage;
