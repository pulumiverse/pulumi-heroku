// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Heroku.Team.Outputs
{

    [OutputType]
    public sealed class GetMembersMemberResult
    {
        public readonly string Email;
        public readonly bool Federated;
        public readonly string Role;
        public readonly string TeamMemberId;
        public readonly bool TwoFactorAuthentication;
        public readonly string UserId;

        [OutputConstructor]
        private GetMembersMemberResult(
            string email,

            bool federated,

            string role,

            string teamMemberId,

            bool twoFactorAuthentication,

            string userId)
        {
            Email = email;
            Federated = federated;
            Role = role;
            TeamMemberId = teamMemberId;
            TwoFactorAuthentication = twoFactorAuthentication;
            UserId = userId;
        }
    }
}
