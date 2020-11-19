#!/usr/bin/env bash

set -eu -o pipefail

pkg=$(dirname "$0")

all_severities=$(awk '
/enum Severity/ { active=1; next }
/}/ { active=0; next }
/^[ \t]*[A-Z][A-Z]*/ { if (active) { print $1 } }
' <$pkg/logpb/log.proto)

severities=()
severities_comments=()
severities_raw_comments=()
for s in $all_severities; do
    raw_comment=$(awk '
/\/\/ '$s'/ { active=1; print $0; next }
/^[ \t]*[A-Z][A-Z]*/ { active=0; next }
{ if (active) { print $0 } }
' <$pkg/logpb/log.proto | sed -e "s/^ *//g")

    if test $s = NONE -o $s = UNKNOWN -o $s = DEFAULT; then
        eval ${s}_COMMENT="\$raw_comment"
        continue
    fi

    comment=$(echo "$raw_comment" | sed -e "s/$s/The $s severity/g")
    set +u # builder image bash is too old and balks at expanding an empty array
    severities=(${severities[*]} $s)
    severities_raw_comments=("${severities_raw_comments[@]}" "$raw_comment")
    severities_comments=("${severities_comments[@]}" "$comment")
    set -u
done

if test $1 = log_channels.go; then
    cat <<EOF
// Code generated by gen.sh. DO NOT EDIT.

package log

import (
  "context"

  "github.com/cockroachdb/cockroach/pkg/util/log/severity"
)

EOF

    for sidx in ${!severities[*]}; do
        SEVERITY=${severities[$sidx]}
        severity=${SEVERITY,,} # FOO -> foo
        severityw=(${severity//_/ }) # foo_bar -> (foo bar)
        Severityt=${severityw[*]^} # (foo bar) -> (Foo Bar)
        Severityt=${Severityt[*]} # (Foo Bar) -> "Foo Bar"
        Severity=${Severityt// /} # "Foo Bar" -> "FooBar"
        SeverityComment=${severities_raw_comments[$sidx]}
        cat <<EOF
// ${Severity}f logs with severity ${SEVERITY},
// if logging has been enabled for the source file where the call is
// performed at the provided verbosity level, via the vmodule setting.
// It extracts log tags from the context and logs them along with the given
// message. Arguments are handled in the manner of fmt.Printf.
//
${SeverityComment}
func ${Severity}f(ctx context.Context, format string, args ...interface{}) {
  logDepth(ctx, 1, severity.${SEVERITY}, format, args...)
}

// V${Severity}f logs with severity ${SEVERITY}.
// It extracts log tags from the context and logs them along with the given
// message. Arguments are handled in the manner of fmt.Printf.
//
${SeverityComment}
func V${Severity}f(ctx context.Context, level Level, format string, args ...interface{}) {
  if VDepth(level, 1) {
    logDepth(ctx, 1, severity.${SEVERITY}, format, args...)
  }
}

// ${Severity} logs with severity ${SEVERITY}.
// It extracts log tags from the context and logs them along with the given
// message.
//
${SeverityComment}
func ${Severity}(ctx context.Context, msg string) {
  logDepth(ctx, 1, severity.${SEVERITY}, msg)
}

// ${Severity}fDepth logs with severity ${SEVERITY},
// offsetting the caller's stack frame by 'depth'.
// It extracts log tags from the context and logs them along with the given
// message. Arguments are handled in the manner of fmt.Printf.
//
${SeverityComment}
func ${Severity}fDepth(ctx context.Context, depth int, format string, args ...interface{}) {
  logDepth(ctx, depth+1, severity.${SEVERITY}, format, args...)
}

EOF
    done
fi

if test $1 = logging.md; then
    echo "# Logging levels (severities)"
    echo

    for sidx in ${!severities[*]}; do
        SEVERITY=${severities[$sidx]}
        SeverityComment=${severities_comments[$sidx]}

        echo "## $SEVERITY"
        echo
        echo "$SeverityComment" | sed -e 's,^// ,,g;s,^// *$,,g'
        echo
    done
fi

if test $1 = severity.go; then
   cat <<EOF
// Code generated by gen.sh. DO NOT EDIT.

package severity

import "github.com/cockroachdb/cockroach/pkg/util/log/logpb"

${UNKNOWN_COMMENT}
const UNKNOWN = logpb.Severity_UNKNOWN

${DEFAULT_COMMENT}
const DEFAULT = logpb.Severity_DEFAULT

${NONE_COMMENT}
const NONE = logpb.Severity_NONE
EOF

   for sidx in ${!severities[*]}; do
       SEVERITY=${severities[$sidx]}
       severity=${SEVERITY,,} # FOO -> foo
       severityw=(${severity//_/ }) # foo_bar -> (foo bar)
       Severityt=${severityw[*]^} # (foo bar) -> (Foo Bar)
       Severityt=${Severityt[*]} # (Foo Bar) -> "Foo Bar"
       Severity=${Severityt// /} # "Foo Bar" -> "FooBar"
       SeverityComment=${severities_raw_comments[$sidx]}
       cat <<EOF

${SeverityComment}
const ${SEVERITY} = logpb.Severity_${SEVERITY}
EOF
   done
fi
