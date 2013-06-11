(defun go-koans-run ()
  "Running Go koans test from go-koans directory and goes to line where error occurs"
  (interactive)
  (setq output (shell-command-to-string "go test "))
  (when (string-match "\\(about_[A-Za-z0-9_]*.go\\):\\([0-9]+\\)" output)

    (let
        ((line (string-to-int (match-string 2 output)))
         (file (expand-file-name (match-string 1 output))))

      (find-file (expand-file-name file))
      (goto-line line)

      (search-forward "__")

      ;;; if expand region exists run it
      ;; if you don't have it https://github.com/magnars/expand-region.el
      (if (fboundp 'er/expand-region)
          (er/expand-region 1))
      )))


(eval-after-load 'go-mode
  '(define-key go-mode-map (kbd "C-c C-r") 'go-koans-run))

(provide 'go-koans)
