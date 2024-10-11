import React, { useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import {
  TextField,
  Button,
  Typography,
  Box,
  Container,
  Paper,
  Snackbar,
  CircularProgress,
} from "@mui/material";
import { Send, ContentCopy } from "@mui/icons-material";
import { shortenUrl } from "../redux/actions";

const ShrinkRayForm = () => {
  const [url, setUrl] = useState("");
  const [snackbarOpen, setSnackbarOpen] = useState(false);
  const dispatch = useDispatch();
  const { shortUrl, loading, error } = useSelector((state) => state.url);

  const handleSubmit = (e) => {
    e.preventDefault();
    dispatch(shortenUrl(url));
  };

  const copyToClipboard = () => {
    navigator.clipboard.writeText(shortUrl);
    setSnackbarOpen(true);
  };

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
          variant="h4"
          component="h1"
          gutterBottom
          align="center"
          sx={{
            color: "primary.main",
            textTransform: "uppercase",
            letterSpacing: 2,
          }}
        >
          ShrinkRay
        </Typography>
        <Box component="form" onSubmit={handleSubmit} sx={{ mt: 2 }}>
          <TextField
            fullWidth
            variant="outlined"
            placeholder="Enter your long URL here"
            value={url}
            onChange={(e) => setUrl(e.target.value)}
            required
            sx={{ mb: 2, input: { color: "text.primary" } }}
          />
          <Button
            type="submit"
            variant="contained"
            color="primary"
            fullWidth
            disabled={loading}
            startIcon={
              loading ? (
                <CircularProgress size={20} color="inherit" />
              ) : (
                <Send />
              )
            }
            sx={{ color: "background.default" }}
          >
            {loading ? "Shrinking..." : "Shrink URL"}
          </Button>
        </Box>
        {error && (
          <Typography color="error" sx={{ mt: 2 }}>
            {error}
          </Typography>
        )}
        {shortUrl && (
          <Box
            sx={{
              mt: 2,
              p: 2,
              backgroundColor: "rgba(255, 232, 31, 0.1)",
              borderRadius: 1,
            }}
          >
            <Typography
              variant="body1"
              sx={{ wordBreak: "break-all", color: "text.primary" }}
            >
              {shortUrl}
            </Typography>
            <Button
              onClick={copyToClipboard}
              variant="outlined"
              color="primary"
              startIcon={<ContentCopy />}
              sx={{ mt: 1, color: "primary.main", borderColor: "primary.main" }}
            >
              Copy
            </Button>
          </Box>
        )}
        <Box sx={{ mt: 4, textAlign: "center" }}>
          <Typography variant="body2" color="text.secondary">
            Shorten, share, and track your links with the power of the Force.
          </Typography>
        </Box>
      </Paper>
      <Snackbar
        open={snackbarOpen}
        autoHideDuration={2000}
        onClose={() => setSnackbarOpen(false)}
        message="Copied to clipboard"
      />
    </Container>
  );
};

export default ShrinkRayForm;
