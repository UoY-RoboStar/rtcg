#!/usr/bin/env sh

usage() {
   echo "usage: $0 [TEMPLATE-DIR] [TRACE-FILE]" >&2
   exit 1
}

main() {
  case $# in
  1)
    template_dir=$1
    trace_file="-" # stdin
    ;;
  2)
    template_dir=$1
    trace_file="$2" # named file
    ;;
  *)
    usage
    ;;
  esac

  ./bin/rtcg-read-traces "${trace_file}" |
    ./bin/rtcg-make-stms - |
    ./bin/rtcg-gen "${template_dir}" -
}

main "$@"